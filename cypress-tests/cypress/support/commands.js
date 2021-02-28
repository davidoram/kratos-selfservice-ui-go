// cypress/support/commands.js

import {
  gen,
  MAIL_API
} from '../helpers'


// delEmails delete *all* emails in Mailslurper
//
// See https://github.com/mailslurper/mailslurper/wiki/Email-Endpoints
Cypress.Commands.add("delEmails", () => {
  console.log("mailslurper delEmails")
  cy.request('DELETE', MAIL_API + '/mail', { pruneCode: 'all' })
    .then((response) => {
      expect(response.status).to.eq(200)
    })
});

// getAllEmails returns the list of emails from Mailslurper
//
// See https://github.com/mailslurper/mailslurper/wiki/Email-Endpoints
Cypress.Commands.add("getAllEmails", () => {
  cy.request('GET', MAIL_API + '/mail')
  .then((response) => {
    expect(response.status).to.eq(200)
    console.log("mailslurper getAllEmails ", response.body.mailItems)
    return response.body.mailItems
  })
});

// pollForEmail polls Mailslurper at least 'retries' times until one or more emails is present.
// It takes some tme for an email sent from kratos to appear in the Mailsurper inbox, so this
// function will poll Mailslurper until an email turns up
//
// See polling technique at https://docs.cypress.io/api/commands/request.html#Request-Polling
function pollForEmail(retries) {
  cy.request('GET', MAIL_API + '/mail')
  .then((response) => {
    if (response.status === 200 && response.body.totalRecords >= 1) {
      return response.body.totalRecords
    }
    if (retries > 0) {
      return pollForEmail(retries - 1)
    }
    return -1
  })
}

// getLatestEmail polls Mailslurper for email 20 times, until at least one email is present in the inbox
// and then returns the last email in the list
//
// See https://github.com/mailslurper/mailslurper/wiki/Email-Endpoints
Cypress.Commands.add("getLatestEmail", () => {
  pollForEmail(20)
  cy.request('GET', MAIL_API + '/mail')
  .then((response) => {
    expect(response.status).to.eq(200)
    console.log("mailslurper getAllEmails ", response.body.mailItems)
    return response.body.mailItems[response.body.totalRecords - 1]
  })
});

// countEmails returns the total number of emails from Mailslurper
//
// See https://github.com/mailslurper/mailslurper/wiki/Email-Endpoints
Cypress.Commands.add("countEmails", () => {
  cy.request('GET', MAIL_API + '/mail')
  .then((response) => {
    expect(response.status).to.eq(200)
    console.log("mailslurper countEmails ", response.body.totalRecords)
    return response.body.totalRecords
  })
});

// mergeFields combines 'form' with new 'fields'
//
const mergeFields = (form, fields) => {
  const result = {}
  form.fields.forEach(({ name, value }) => {
    result[name] = value
  })

  return { ...result, ...fields }
}


// registerApi creates a new user in Kratos.
//
// Its a fast way of creating a user for a test, and set the 'user' alias:
//  {
//    email: "{random email}",
//    password: "{random password}"
//  }
//
Cypress.Commands.add('registerApi',({ email = gen.email(), password = gen.password(), fields = {} } = {} ) =>
  cy.request({
      url: '/self-service/registration/api'
  }).then(({ body }) => {
    const form = body.methods.password.config
    return cy.request({
      method: form.method,
      body: mergeFields(form, {
        ...fields,
        'traits.email': email,
        password
      }),
      url: form.action
    })
  }).then(({ body }) => {
    expect(body.identity.traits.email).to.contain(email)
  }).then( () => {
    return cy.wrap({email: email, password: password}).as('user')
  })
);
