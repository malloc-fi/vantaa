'use strict';

var webpack = require('webpack'),
    HtmlWebpackPlugin = require('html-webpack-plugin'),
    path = require('path'),
    srcPath = path.join(__dirname, 'src');

module.exports = {
  target: 'web',
  cache: true,
  entry: {
    module: path.join(srcPath, 'module.js'),
    common: ['react', 'react-router', 'events']
  },
  resolve: {
    root: srcPath,
    extensions: ['', '.js'],
    modulesDirectories: ['node_modules', 'src']
  },
  output: {
    path: path.join(__dirname, 'tmp'),
    publicPath: '',
    filename: '[name].js',
    library: ['VantaaAdmin', '[name]'],
    pathInfo: true
  },

  module: {
    loaders: [
      {test: /\.js?$/, exclude: /node_modules/, loader: 'babel-loader'},
      {test: /\.css$/, exclude: /node_modules/, loader: 'style!css?sourceMap!postcss'}
    ]
  },

  postcss: [
    require('cssnext'),
    require('lost'),
    require('autoprefixer'),
    require('precss')
  ],

  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new webpack.optimize.CommonsChunkPlugin('common', 'common.js'),
    new HtmlWebpackPlugin({
      inject: true,
      template: 'src/index.html'
    }),
    new webpack.NoErrorsPlugin(),
    new webpack.IgnorePlugin(/^\.\/locale$/, /moment$/)
  ],

  debug: true,
  devtool: 'eval-cheap-module-source-map',
  devServer: {
    contentBase: './tmp',
    historyApiFallback: true
  }
};
