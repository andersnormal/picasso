// theme.config.js
export default {
    github: 'https://github.com/andersnormal/picasso',
    docsRepositoryBase: 'https://github.com/andersnormal/picasso/edit/main/docs/pages', // base URL for the docs repository
    titleSuffix: ' – Picasso',
    nextLinks: true,
    prevLinks: true,
    search: true,
    customSearch: null, // customizable, you can use algolia for example
    darkMode: true,
    footer: true,
    footerText: `MIT ${new Date().getFullYear()} © Sebastian Doell (@katallaxie).`,
    footerEditLink: `Edit this page on GitHub`,
    logo: (
      <>
        <svg>...</svg>
        <span>Picasso</span>
      </>
    )
  }