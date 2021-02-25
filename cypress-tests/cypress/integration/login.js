/// <reference types="cypress" />

import * as shared from './shared.js'

context('Login', () => {
  var user = {}
  beforeEach(() => {
    cy.visit('/auth/logout')
    user = shared.registerAndLogin()
    cy.visit('/auth/logout')
  })

  it('allows a user to login', () => {
    // Navigate to the registration page
    cy.visit('/auth/login')
    cy.get('[data-cy=page-heading]').should('contain', 'Login')
    cy.get('[data-cy=password]').type(user.password)
    cy.get("[data-cy='identifier']").type(user.email)
    cy.get('[data-cy=submit]').click()

    // Should be redirected to sucess page
    cy.get('[data-cy=flash_info]').should('contain', 'Logged in')

    // Navigate to protected page - ok
    cy.visit('/dashboard')
    cy.get('[data-cy=page-heading]').should('contain', 'Dashboard')

    // Logout
    cy.visit('/auth/logout')

    // Navigate to protected page - redirect to login
    cy.visit('/dashboard')
    cy.get('[data-cy=page-heading]').should('contain', 'Login')

  })

  it('fails login with invalid creds', () => {
    // Navigate to the registration page
    cy.visit('/auth/login')
    cy.get('[data-cy=page-heading]').should('contain', 'Login')
    cy.get('[data-cy=password]').type("wrong password")
    cy.get("[data-cy='identifier']").type(user.email)
    cy.get('[data-cy=submit]').click()

    // Should remain on the login page
    cy.get('[data-cy=page-heading]').should('contain', 'Login')

    // Navigate to protected page - redirect to login
    cy.visit('/dashboard')
    cy.get('[data-cy=page-heading]').should('contain', 'Login')

  })

})
