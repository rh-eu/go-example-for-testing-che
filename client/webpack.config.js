const webpack = require('webpack'); 

const HtmlWebPackPlugin = require("html-webpack-plugin");

const path = require('path');

const BUILD_DIR = path.resolve(__dirname, '../sitedata/built');

module.exports = {
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader"
        }
      },
      {
        test: /\.html$/,
        use: [
          {
            loader: "html-loader"
          }
        ]
      }
    ]
  },
  output: {
    path: BUILD_DIR,
    publicPath: '/built/',
    filename: 'bundle.js'
  },
  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new HtmlWebPackPlugin({
      template: "../sitedata/index.html",
      filename: "index.html"
    })
  ],
  devServer: {
    contentBase: "../sitedata",
    inline: true,
    hot: true,
    host: "0.0.0.0",
    port: 4100
    //proxy: {
    //  "**": "http://localhost:5051"
    //}
  }
};