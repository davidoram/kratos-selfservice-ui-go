// cypress/support/commands.js

import {
  APP_URL, gen
} from '../helpers'

const { MailSlurp } = require("mailslurp-client");

const mailslurp = new MailSlurp({apiKey: Cypress.env('MAILSLURP_API_KEY')});

Cypress.Commands.add("createInbox", ({ email = gen.email() }) => {
  return cy.wrap(mailslurp.createInbox({emailAddress: email})).as('inbox');
});

Cypress.Commands.add("waitForLatestEmail", (inboxId) => {
  return mailslurp.waitForLatestEmail(inboxId);
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
