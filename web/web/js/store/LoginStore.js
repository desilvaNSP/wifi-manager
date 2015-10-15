/*
 * Copyright (c) 2014, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 *
 * TodoStore
 */

var AppDispatcher = require('../dispatcher/AppDispatcher');
var EventEmitter = require('events').EventEmitter;
var DashBoardConstants = require('../constants/DashBoardConstants');
var assign = require('object-assign');

var CHANGE_EVENT = 'change';

var logedIn = false;

/**
 * Create a TODO item.
 * @param  {string} text The content of the TODO
 */
function login(username,password) {
  console.log("<<<<<<< loggingin" + username)
  logedIn = true;
}

/**
 * Update a TODO item.
 * @param  {string} id
 * @param {object} updates An object literal containing only the data to be
 *     updated.
 */
function update(id, updates) {
  logedIn = true;
}

var LoginStore = assign({}, EventEmitter.prototype, {

  /**
   * Get the entire collection of TODOs.
   * @return {object}
   */
  getLoginState: function() {
    console.log("geting login state " + logedIn)
    return logedIn;
  },

  emitChange: function() {
    this.emit(CHANGE_EVENT);
  },

  /**
   * @param {function} callback
   */
  addChangeListener: function(callback) {
    this.on(CHANGE_EVENT, callback);
  },

  /**
   * @param {function} callback
   */
  removeChangeListener: function(callback) {
    this.removeListener(CHANGE_EVENT, callback);
  }
});

// Register callback to handle all updates
AppDispatcher.register(function(action) {
  switch(action.actionType) {
    case DashBoardConstants.LOGIN:
        login(action.username,action.password);
        LoginStore.emitChange();
      break;

    default:
      // no op
  }
});

module.exports = LoginStore;
