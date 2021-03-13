/// <reference types="cypress" />

context('Registration', () => {

  function makeUser() {
    const uuid = () => Cypress._.random(0, 1e12)
    // Fill out details for a new user
    var user = {
      password: "abc123Pass#",
      email: "bob" + uuid() + "@gmail.com",
      firstname: "Bob",
      lastname: "Smith"
    };
    return user
  }

  it('allows a user to register sucesfully', () => {

    cy.visit('/')

    // Navigate to the registration page
    cy.get('[data-cy=registration]').first().click()
    cy.get('[data-cy=page-heading]').should('contain', 'Registration')

    // Fill out details for a new user
    var user = makeUser()
    cy.get('[data-cy=password]').type(user.password)
    cy.get("[data-cy='traits.email']").type(user.email)
    cy.get("[data-cy='traits.name.first']").type(user.firstname)
    cy.get("[data-cy='traits.name.last']").type(user.lastname)
    cy.get('[data-cy=submit]').click()

    // Should be redirected to sucess page
    cy.get('[data-cy=flash_info]').should('contain', 'Registration complete')
  })

  it('detects error during registration, password not secure enough', () => {
    cy.visit('/')

    // Navigate to the registration page
    cy.get('[data-cy=registration]').first().click()

    cy.get('[data-cy=page-heading]').should('contain', 'Registration')

    // Fill out details for a new user
    var user = makeUser()
    // use an insecure password
    user.password = "password"
    cy.get('[data-cy=password]').type(user.password)
    cy.get("[data-cy='traits.email']").type(user.email)
    cy.get("[data-cy='traits.name.first']").type(user.firstname)
    cy.get("[data-cy='traits.name.last']").type(user.lastname)
    cy.get('[data-cy=submit]').click()

    // Should stay on the same page
    cy.get('[data-cy=page-heading]').should('contain', 'Registration')

    // Should display an error against the password field
    cy.get("[data-cy='field_message_id_password']").should('contain', '4000005')

    // Fix the error, should accept ok
    cy.get('[data-cy=password]').type("abc123Pass#")

    // Sumbit the form again
    cy.get('[data-cy=submit]').click()

    // Should be redirected to sucess page
    cy.get('[data-cy=flash_info]').should('contain', 'Registration complete')
  })
})