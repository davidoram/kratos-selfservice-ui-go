/// <reference types="cypress" />

context('Login', () => {

  beforeEach(() => {
    cy.registerApi()
  })

  it('allows a user to login', function () {
    // Navigate to the registration page
    cy.visit('/auth/login')
    cy.get('[data-cy=page-heading]').should('contain', 'Login')
    cy.get('[data-cy=password]').type(this.user.password)
    cy.get("[data-cy='identifier']").type(this.user.email)
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

  it('fails login with invalid creds', function () {
    // Navigate to the registration page
    cy.visit('/auth/login')
    cy.get('[data-cy=page-heading]').should('contain', 'Login')
    cy.get('[data-cy=password]').type("wrong password")
    cy.get("[data-cy='identifier']").type(this.user.email)
    cy.get('[data-cy=submit]').click()

    // Should remain on the login page
    cy.get('[data-cy=page-heading]').should('contain', 'Login')

    // Navigate to protected page - redirect to login
    cy.visit('/dashboard')
    cy.get('[data-cy=page-heading]').should('contain', 'Login')

  })

})
