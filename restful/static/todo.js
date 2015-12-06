/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function ReminderCtrl($scope, $http) {
  $scope.reminders = [];
  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/api/reminders').
          success(function(data) { $scope.reminders = data; console.log(data);}).
      error(logError);
  };

  $scope.addTodo = function() {
    $scope.working = true;
    $http.post('/api/reminders/', {Message: $scope.todoText}).
      error(logError).
      success(function() {
        refresh().then(function() {
          $scope.working = false;
          $scope.todoText = '';
        });
      });
  };

  $scope.toggleDone = function(reminder) {
      data = {ID: reminder.ID, Message: reminder.Message};
    $http.put('/api/reminders/'+reminder.ID, data).
      error(logError).
          success(function() { reminder.Done = !reminder.Done; });
  };

  refresh().then(function() { $scope.working = false; });
}
