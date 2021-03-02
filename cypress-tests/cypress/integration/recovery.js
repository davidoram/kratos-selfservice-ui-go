/// <reference types="cypress" />

import {
  gen,
} from '../helpers'

describe('Recovery', () => {

  beforeEach(() => {
    cy.registerApi()
    .then(() => {
      // Delete the account verification email
      cy.mhDeleteAll();
    })
  })

  it('sends a recovery email', function () {
    // Navigate to the recovery page

    cy.visit('/auth/recovery')
    cy.get('[data-cy=email]').type(this.user.email)
    cy.get('[data-cy=submit]').click()
    console.log("clicked");
    // State should be updated
    cy.get('[data-cy=state]').should('have.attr', 'data-value', 'sent_email')
    .then( () => {
      cy.waitUntil(() =>
        cy.mhGetMailsMatchSubject('.*Recover_access_to_your_account.*')
        .then( mails => mails.length > 0))

    })
    .then( () => {
      cy.mhGetMailsMatchSubject('.*Recover_access_to_your_account.*')
      .mhFirst()
      .mhGetBody()
      .then ( (body) => {
        // Extract the confirmation link
        cy.mhGetLink(body)
        .then( (url) => {
          console.log("url: ", url);
          // Visit the link
          cy.visit(url);
        })
      })
    })
  })

})
