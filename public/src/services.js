'use strict';
define(['angular'], function(angular) {
  /* Services */
  angular.module('app.services', [])
    .factory('Config', [function() {
      var Config = {
        url: "http://localhost:8080/api/"
      };
      return Config;
    }])
    // with users
    .factory('UserService', ['$http', 'Config', function($http, Config) {
      var User = {
        type: 'users',
        userOBJ: {},
        username: "",
        password: "",
        token: "",
        setUsername: function(inputUsername) {
          User.username = inputUsername;
        },

        getUserName: function() {
          return this.username;
        },

        setToken: function(token) {
          User.token = token;
        },

        getToken: function() {
          return this.token;
        },

        getLoggedUserName: function() {
          return localStorage.getItem("username");
        },

        setPassword: function(inputPassword) {
          User.password = inputPassword;
        },
        getPassword: function() {
          if (this.password != "") {
            return this.password;
          } else {
            return localStorage.getItem("password");
          }
        },
        storeUserLocally: function(obj) {
          if (typeof(Storage) !== "undefined") {
            localStorage.setItem("username", this.getUserName());
            localStorage.setItem("password", this.getPassword());
            localStorage.setItem("token", this.getToken());
          } else {
            console.log('no local storage available');
          }
        },

        getUserOBJ: function() {
          return $this.userOBJ;
        },
        login: function(inputUsername, inputPassword, callback) {
          var headers = {
            'Content-Type': 'application/json'
          };
          var jsonObject = angular.toJson({
            "username": inputUsername,
            "password": inputPassword
          });
          $http.post(Config.url + this.type + '/session', jsonObject, {
              headers: headers
            })
            .then(function(response) {
              User.setUsername(response.user.username);
              User.setToken(response.token);
              User.storeUserLocally(response.user);
              callback(response);
            })
            .catch(function(err) {
              callback(err);
            })
        },
        getUsers: function(page = 0, size = 10, success, error) {
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.get(Config.url + this.type, {
              params: {
                "page": page,
                "size": size
              }
            }, {
              headers: headers
            })
            .then(function(response) {
              success(response);
            })
            .catch(function(err) {
              error(err);
            })
        },
        logout: function() {

        }
      }
      return User;
    }])
    // with orders
    .factory('OrderService', ['$http', 'Config', function($http, Config) {
      var Order = {
        type: 'orders',
        getOrders: function(page = 0, size = 10, catalog = 1, success, error) {
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.get(Config.url + this.type, { params: { "page": page, "size": size, "catalog": catalog} }, {headers: headers})
            .then(function(response) {
              success(response);
            })
            .catch(function(err) {
              error(err);
            })
        }
      }
      return Order;
    }])
});
