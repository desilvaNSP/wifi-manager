/*
 * Copyright (c) 2014-2015, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 *
 * TodoActions
 */

var AppDispatcher = require('../dispatcher/AppDispatcher');
var DashBoardConstants = require('../constants/DashBoardConstants');

var DashBoardActions = {

  /**
   * @param  {string} text
   */
  login: function(username, password) {
    console.log("dispatching" + username)
    AppDispatcher.dispatch({
      actionType: DashBoardConstants.LOGIN,
      username: username,
      password: password
    });
  },

 
};

module.exports = DashBoardActions;
