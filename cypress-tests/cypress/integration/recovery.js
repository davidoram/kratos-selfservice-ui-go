/// <reference types="cypress" />

import {
  gen,
} from '../helpers'

describe('Recovery', () => {

  beforeEach(() => {
    cy.registerApi()
    .then(() => {
      // Delete the account verification email
      cy.delEmails()
    })
  })

  it('sends a recovery email', function () {
    // Navigate to the recovery page

    cy.visit('/auth/recovery')
    cy.get('[data-cy=email]').type(this.user.email)
    cy.get('[data-cy=submit]').click()

    // State should be updated
    cy.get('[data-cy=state]').should('have.attr', 'data-value', 'sent_email')

    // Get the recovery email
    cy.getAllEmails().then( function() {
      cy.log("Email subject:", this.emails[0].subject)
    }).then(function() {
      // verify that email contains the code
      assert.strictEqual(/please verify your account by clicking the following link/.test(this.emails[0].body), true);
    }).then(function() {
      // extract the link & click on it
      var links = this.emails[0].body.match(/href=\"(.*)\" rel/);
      assert.isDefined(links);
      expect(links.length).to.equal(2)
      return links[1]
    }).then(function( link ) {
      cy.visit(link)
    })
  })

})
