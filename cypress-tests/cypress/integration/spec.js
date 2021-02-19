it('can see the homepage', () => {
  cy.visit('/')

  cy.get('h1')
    .should('contain', 'Homepage')
})