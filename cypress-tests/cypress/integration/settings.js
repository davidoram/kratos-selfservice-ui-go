/// <reference types="cypress" />

context('Settings', () => {
  beforeEach(() => {
    cy.registerAndLogin()
  })


  it('allows a user to update their profile', () => {
    // Navigate to the registration page
    cy.visit('/auth/settings')
    cy.get('[data-cy=page-heading]').should('contain', 'Update Profile')
    cy.get("[data-cy='profile_traits.name.first']").clear()
    cy.get("[data-cy='profile_traits.name.first']").type("Robert")
    cy.get("[data-cy='profile_traits.name.last']").clear()
    cy.get("[data-cy='profile_traits.name.last']").type("Smitty")
    cy.get('[data-cy=profile_submit]').click()

    // Should be redirected to sucess page
    cy.get('[data-cy=flash_info]').should('contain', 'Settings updated')

    // Reload and check details saved
    cy.visit('/auth/settings')
    cy.get("[data-cy='profile_traits.name.first']").should('have.value', 'Robert')
    cy.get("[data-cy='profile_traits.name.last']").should('have.value', 'Smitty')
  })

   it('allows a user to update their password', function() {
    var newPassword = "ghg65svsbs%";
    // Navigate to the registration page
    cy.visit('/auth/settings')
    cy.get('[data-cy=page-heading]').should('contain', 'Update Profile')
    cy.get("[data-cy='password_password']").clear()
    cy.get("[data-cy='password_password']").type(newPassword)
    cy.get('[data-cy=password_submit]').click()

    // Should be redirected to sucess page
    cy.get('[data-cy=flash_info]').should('contain', 'Settings updated')

    // Logout and verify password has changed
    cy.visit('/auth/logout')
    cy.visit('/auth/login')
    cy.get('[data-cy=password]').type(newPassword)
    cy.get("[data-cy='identifier']").type(this.user.email)
    cy.get('[data-cy=submit]').click()

    // Should be redirected to sucess page
    cy.get('[data-cy=flash_info]').should('contain', 'Logged in')

  })


})
