/// <reference types="cypress" />

import {
  gen,
} from '../helpers'

describe('Recovery', () => {
  beforeEach(() => {
    cy.delEmails()
    cy.registerApi()
    // Wait for account verification email
    cy.getLatestEmail().then((email) => {
      console.log("setup email", email)
      // Delete it
      cy.delEmails()
    })
  })

  // it('allows a user to login', function () {
  //   // Navigate to the login page
  //   console.log(this.user);
  //   cy.visit('/auth/login')
  //   cy.get('[data-cy=password]').type(this.user.password)
  //   cy.get("[data-cy='identifier']").type(this.user.email)
  //   cy.get('[data-cy=submit]').click()

  //   // Should be redirected to sucess page
  //   cy.get('[data-cy=flash_info]').should('contain', 'Logged in')
  // })

  it('sends a recovery email', function () {
    // Navigate to the recovery page
    console.log('-------------------------------');
    console.log(this.user);

    cy.visit('/auth/recovery')
    cy.get('[data-cy=email]').type(this.user.email)
    cy.get('[data-cy=submit]').click()

    // State should be updated
    cy.get('[data-cy=state]').should('have.attr', 'data-value', 'sent_email')

    cy.countEmails().then((count) => {
      cy.log("Email count", count)
    })
    // Get the recovery email
    cy.getLatestEmail().then((email) => {
      // verify we received an email
      assert.isDefined(email);
      cy.log(email.subject)

      // verify that email contains the code
      assert.strictEqual(/please verify your account by clicking the following link/.test(email.body), true);

      // extract the link
      var links = email.body.match(/href=\"(.*)\" rel/);
      console.log(links);
      assert.isDefined(links);
      expect(links.length).to.equal(2)
      console.log(links[1])
      cy.visit(links[1])
    });
  })

})
