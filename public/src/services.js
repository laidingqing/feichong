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
          console.log(obj)
          if (typeof(Storage) !== "undefined") {
            localStorage.setItem("username", obj.name);
            localStorage.setItem("token", obj.sessionKey);
          } else {
            console.log('no local storage available');
          }
        },

        getUserOBJ: function() {
          return $this.userOBJ;
        },
        session: function(data, callback, error) {
          var jsonObject = angular.toJson(data);
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.post(Config.url + 'session/', jsonObject, {
              headers: headers
            })
            .then(function(response) {
              console.log(response.data)
              User.storeUserLocally(response.data);
              callback(response);
            })
            .catch(function(err) {
              console.log(err)
              error(err);
            })
        },
        putUser: function(user, callback) {
          var headers = {
            'Content-Type': 'application/json'
          };
          var jsonObject = angular.toJson(user);
          $http.put(Config.url + this.type + '/' + user.id +"/", jsonObject, {
              headers: headers
            })
            .then(function(response) {
              console.log(response)
              callback(response);
            })
            .catch(function(err) {
              callback(err);
            })
        },
        getUser: function(userId, success, error) {
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.get(Config.url + this.type + "/" + userId + "/", {}, {
              headers: headers
            })
            .then(function(response) {
              success(response.data);
            })
            .catch(function(err) {
              error(err);
            })
        },
        checkUserName: function(username, success, error) {
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.get(Config.url + "/checkname/", { params: { "username": username} }, {
              headers: headers
            })
            .then(function(response) {
              success(response.data);
            })
            .catch(function(err) {
              error(err);
            })
        },
        putUserSecurity: function(user, callback) {
          var headers = {
            'Content-Type': 'application/json'
          };
          var jsonObject = angular.toJson(user);
          $http.put(Config.url + this.type + '/' + user.id +"/security/", jsonObject, {
              headers: headers
            })
            .then(function(response) {
              console.log(response)
              callback(response);
            })
            .catch(function(err) {
              callback(err);
            })
        },
        getUserBySelf: function(success, error) {
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.get(Config.url + "operators", {}, {
              headers: headers
            })
            .then(function(response) {
              success(response);
            })
            .catch(function(err) {
              error(err);
            })
        },
        getUserByEnterprise: function(success, error) {
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.get(Config.url + "companies/", {}, {
              headers: headers
            })
            .then(function(response) {
              success(response);
            })
            .catch(function(err) {
              error(err);
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
        getOrders: function(page = 0, size = 10, orderNO, catalog = 1, success, error) {
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.get(Config.url + this.type, { params: { "page": page, "size": size, "catalog": catalog, "orderNO": orderNO} }, {headers: headers})
            .then(function(response) {
              success(response);
            })
            .catch(function(err) {
              error(err);
            })
        },

        getBusinessByOrder: function(orderId, success, error) {
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.get(Config.url + this.type + "/" + orderId + "/business/", { params: { } }, {headers: headers})
            .then(function(response) {
              success(response.data);
            })
            .catch(function(err) {
              error(err);
            })
        },

        getBusinessByID: function(orderId, businessId, success, error) {
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.get(Config.url + this.type + "/" + orderId + "/business/" + businessId + "/", { params: { } }, {headers: headers})
            .then(function(response) {
              success(response.data);
            })
            .catch(function(err) {
              error(err);
            })
        },

        postOrder: function(order, callback) {
          var headers = {
            'Content-Type': 'application/json'
          };
          var jsonObject = angular.toJson(order);
          $http.post(Config.url + this.type + '/', jsonObject, {
              headers: headers
            })
            .then(function(response) {
              console.log(response)
              callback(response);
            })
            .catch(function(err) {
              callback(err);
            })
        },
      }
      return Order;
    }])
    .factory('BusinessService', ['$http', 'Config', function($http, Config) {
      var Business = {
        type: 'business',
        createBusiness: function(orderId, data, callback, error){
          var jsonObject = angular.toJson(data);
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.post(Config.url + 'orders/' + orderId +"/business/", jsonObject, {
              headers: headers
            })
            .then(function(response) {
              console.log(response)
              callback(response);
            })
            .catch(function(err) {
              error(err);
            })
        },
        removeBusiness: function(orderId, businessId, callback, error){
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.delete(Config.url + 'orders/' + orderId +"/business/" + businessId + "/", {}, {
              headers: headers
            })
            .then(function(response) {
              console.log(response)
              callback(response);
            })
            .catch(function(err) {
              error(err);
            })
        },
        putCapitalInfo: function(businessId, data, callback, error){
          var jsonObject = angular.toJson(data);
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.put(Config.url + this.type + '/' + businessId +"/capitals/", jsonObject, {
              headers: headers
            })
            .then(function(response) {
              console.log(response)
              callback(response);
            })
            .catch(function(err) {
              error(err);
            })
        },
        putProfitInfo: function(businessId, data, callback, error){
          var jsonObject = angular.toJson(data);
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.put(Config.url + this.type + '/' + businessId +"/profits/", jsonObject, {
              headers: headers
            })
            .then(function(response) {
              console.log(response)
              callback(response);
            })
            .catch(function(err) {
              error(err);
            })
        },
        putTaxInfo: function(businessId, data, callback, error){
          var jsonObject = angular.toJson(data);
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.put(Config.url + this.type + '/' + businessId +"/tax/", jsonObject, {
              headers: headers
            })
            .then(function(response) {
              console.log(response)
              callback(response);
            })
            .catch(function(err) {
              error(err);
            })
        }
      }
      return Business;
    }])

    .factory('FeedbackService', ['$http', 'Config', function($http, Config) {
      var Feedback = {
        type: 'consults',
        getFeedbacks: function(page = 0, size = 10, success, error){
          var headers = {
            'Content-Type': 'application/json'
          };
          $http.get(Config.url + this.type, { params: { "page": page, "size": size} }, {headers: headers})
            .then(function(response) {
              success(response);
            })
            .catch(function(err) {
              error(err);
            })
        }
      }
      return Feedback
    }])
})
