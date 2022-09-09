const SpeedMeasurePlugin = require('speed-measure-webpack-plugin');
const sassResourcesLoader = require('craco-sass-resources-loader');

const smp = new SpeedMeasurePlugin();

module.exports = {
  plugins: [
    {
      plugin: sassResourcesLoader,
      options: {
        resources: [
          './src/stylesheets/_uswdsUtilities.scss',
          './src/stylesheets/_colors.scss',
          './src/stylesheets/_variables.module.scss'
        ]
      }
    },
    {
      plugin: {
        overrideWebpackConfig: ({ webpackConfig }) => {
          return smp.wrap(webpackConfig);
        }
      }
    }
  ],
  typescript: {
    enableTypeChecking: false
  }
};
