/// <reference types="cypress" />

import * as shared from './shared.js'

context('Settings', () => {
  var user = {}
  beforeEach(() => {
    cy.visit('/auth/logout')
    user = shared.registerAndLogin()
  })


  it('allows a user to update their profile', () => {
    // Navigate to the registration page
    cy.visit('/auth/settings')
    cy.get('[data-cy=page-heading]').should('contain', 'Update Profile')
    cy.get("[data-cy='traits.name.first']").clear()
    cy.get("[data-cy='traits.name.first']").type("Robert")
    cy.get("[data-cy='traits.name.last']").clear()
    cy.get("[data-cy='traits.name.last']").type("Smitty")
    cy.get('[data-cy=submit]').click()

    // Should be redirected to sucess page
    cy.get('[data-cy=flash_info]').should('contain', 'Settings updated')

    // Reload and check details saved
    cy.visit('/auth/settings')
    cy.get("[data-cy='traits.name.first']").should('have.value', 'Robert')
    cy.get("[data-cy='traits.name.last']").should('have.value', 'Smitty')


  })


})
