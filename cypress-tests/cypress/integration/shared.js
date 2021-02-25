
export function registerAndLogin() {
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
