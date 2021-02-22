/// <reference types="cypress" />

context('Login', () => {
  var user = {}
  beforeEach(() => {
    cy.visit('/auth/logout')
    user = registerUser()
    cy.visit('/auth/logout')
  })

  function registerUser() {
    const uuid = () => Cypress._.random(0, 1e12)

    // Navigate to the registration page
    cy.visit('/auth/registration')
    cy.get('[data-cy=page-heading]').should('contain', 'Registration')

    // Fill out details for a new user
    var user = {
      password: "abc123Pass#",
      email: "bob" + uuid() + "@gmail.com",
      firstname: "Bob",
      lastname: "Smith"
    };

    cy.get('[data-cy=password]').type(user.password)
    cy.get("[data-cy='traits.email']").type(user.email)
    cy.get("[data-cy='traits.name.first']").type(user.firstname)
    cy.get("[data-cy='traits.name.last']").type(user.lastname)
    cy.get('[data-cy=submit]').click()

    // Should be redirected to sucess page
    cy.get('[data-cy=flash_info]').should('contain', 'Registration complete')

    return user
  }

  it('allows a user to login', () => {
    // Navigate to the registration page
    cy.visit('/auth/login')
    cy.get('[data-cy=page-heading]').should('contain', 'Login')
    cy.get('[data-cy=password]').type(user.password)
    cy.get("[data-cy='identifier']").type(user.email)
    cy.get('[data-cy=submit]').click()

    // Should be redirected to sucess page
    cy.get('[data-cy=flash_info]').should('contain', 'Logged in')
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
  })

})
