it('detects error during registration, password not secure enough', () => {
  const uuid = () => Cypress._.random(0, 1e12)

  cy.visit('/')

  // Navigate to the registration page
  cy.get('[data-cy=registration]').click()
  cy.get('[data-cy=page-heading]').should('contain', 'Registration')

  // Fill out details for a new user
  var user = {
    password: "password",
    email: "bob" + uuid() + "@gmail.com",
    firstname: "Bob",
    lastname: "Smith"
  };

  cy.get('[data-cy=password]').type(user.password)
  cy.get("[data-cy='traits.email']").type(user.email)
  cy.get("[data-cy='traits.name.first']").type(user.firstname)
  cy.get("[data-cy='traits.name.last']").type(user.lastname)
  cy.get('[data-cy=submit]').click()

  // Should stay on the same page
  cy.get('[data-cy=page-heading]').should('contain', 'Registration')

  // Should display an error against the password field
  cy.get("[data-cy='field_message_id_password']").should('contain', '4000005')
})