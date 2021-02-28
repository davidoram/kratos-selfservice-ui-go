// cypress/support/commands.js

import {
  gen,
  MAIL_API
} from '../helpers'


// See https://github.com/mailslurper/mailslurper/wiki/Email-Endpoints
Cypress.Commands.add("delEmails", () => {
  console.log("mailslurper delEmails")
  cy.request('DELETE', MAIL_API + '/mail', { pruneCode: 'all' })
    .then((response) => {
      expect(response.status).to.eq(200)
    })
});

Cypress.Commands.add("getAllEmails", () => {
  cy.request('GET', MAIL_API + '/mail')
  .then((response) => {
    expect(response.status).to.eq(200)
    console.log("mailslurper getAllEmails ", response.body.mailItems)
    return response.body.mailItems
  })
});

function pollForEmail(retries) {
  // See polling technique at https://docs.cypress.io/api/commands/request.html#Request-Polling
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

Cypress.Commands.add("getLatestEmail", () => {
  pollForEmail(20)
  cy.request('GET', MAIL_API + '/mail')
  .then((response) => {
    expect(response.status).to.eq(200)
    console.log("mailslurper getAllEmails ", response.body.mailItems)
    return response.body.mailItems[response.body.totalRecords - 1]
  })
});

Cypress.Commands.add("countEmails", () => {
  cy.request('GET', MAIL_API + '/mail')
  .then((response) => {
    expect(response.status).to.eq(200)
    console.log("mailslurper countEmails ", response.body.totalRecords)
    return response.body.totalRecords
  })
});

const mergeFields = (form, fields) => {
  const result = {}
  form.fields.forEach(({ name, value }) => {
    result[name] = value
  })

  return { ...result, ...fields }
}


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
