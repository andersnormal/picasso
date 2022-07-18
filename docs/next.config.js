const isProd = process.env.NODE_ENV === "production";

const withNextra = require('nextra')({
    theme: 'nextra-theme-docs',
    themeConfig: './theme.config.js',
    unstable_staticImage: true,
})
module.exports = withNextra({ 
    basePath: isProd ? "/picasso/" : "",
    assetPrefix: isProd ? "/picasso/" : "", experimental: {
    images: {
        unoptimized: true
    }
}})