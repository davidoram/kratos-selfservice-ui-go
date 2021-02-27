/// <reference types="cypress" />

import {
  gen,
} from '../helpers'

describe('Recovery', () => {
  beforeEach(() => {
    var user = gen.identity()
    cy.registerApi({email: user.email, password: user.password})
    cy.createInbox({email: user.email})
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
    console.log(this.user);
    console.log(this.inbox);
    cy.visit('/auth/recovery')
    cy.get('[data-cy=email]').type(this.user.email)
    cy.get('[data-cy=submit]').click()

    // State should be updated
    cy.get('[data-cy=state]').should('have.attr', 'data-value', 'sent_email')
  })

})
