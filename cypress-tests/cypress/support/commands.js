// cypress/support/commands.js

import {
  gen,
  MAIL_API
} from '../helpers'


// delEmails delete *all* emails in Mailslurper
//
// See https://github.com/mailslurper/mailslurper/wiki/Email-Endpoints
Cypress.Commands.add("delEmails", () => {
  cy.log("mailslurper delEmails")
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
    return response.body.mailItems
  }).then (( mailItems) => {
    return cy.wrap(mailItems).as('emails')
  })
});


// countEmails returns the total number of emails from Mailslurper
//
// See https://github.com/mailslurper/mailslurper/wiki/Email-Endpoints
// Cypress.Commands.add("countEmails", () => {
//   cy.request('GET', MAIL_API + '/mail')
//   .then((response) => {
//     expect(response.status).to.eq(200)
//     cy.log("mailslurper countEmails ", response.body.totalRecords)
//     return response.body.totalRecords
//   })
// });

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
Cypress.Commands.add('registerApi',({ email = gen.email(), password = gen.password(), fields = {} } = {} ) => {
  cy.clearCookies();
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
});

// registerAndLogin logs out, creates a new user in Kratos, and logs them in
//
// Its a fast way of creating a user for a test, and set the 'user' alias:
//  {
//    email: "{random email}",
//    password: "{random password}"
//  }
//
Cypress.Commands.add('registerAndLogin',() => {
  cy.visit('/auth/logout').then( () => {
    cy.registerApi().then( function() {
      cy.visit('/auth/login')
      cy.get('[data-cy=page-heading]').should('contain', 'Login')
      cy.get('[data-cy=password]').type(this.user.password)
      cy.get("[data-cy='identifier']").type(this.user.email)
      cy.get('[data-cy=submit]').click()

      // Should be redirected to sucess page
      cy.get('[data-cy=flash_info]').should('contain', 'Logged in')
    })
  })
})
