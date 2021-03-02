// cypress/support/commands.js

import {
  gen,
  MAIL_API
} from '../helpers'

// mhGetMailsMatchSubject returns the email where the subject matches the
// regex suppied
//
Cypress.Commands.add('mhGetMailsMatchSubject', (regexStr) => {
  var regex = new RegExp(regexStr);
  console.log("regex:", regex);
  cy.mhGetAllMails().then((mails) => {
    return mails.filter((mail) => {
      console.log("sub: '" + mail.Content.Headers.Subject[0] + "'");
      return regex.test(mail.Content.Headers.Subject[0])
    });
  });
});

// mhGetLink returns the first link from the body of the email as a URL
//
Cypress.Commands.add('mhGetLink', (body) => {
  console.log("body: ", body);
  let re = /"(http[\s\S]*?)">/g;
  let links = re.exec(body);
  if (links != null) {
    return links[1];
  }
})


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
