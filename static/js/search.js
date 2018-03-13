(function(window, document, undefined) {
  "use strict";

  var SPINNER = '<i class="fa fa-2x fa-pulse fa-spinner"></i>';
  var NO_RESULTS = '<span>No results found</span>';

  window.onload = function() {
    var form = document.getElementById("search-form");
    form.addEventListener("submit", handleFormSubmit);
  };

  var searchContainer = document.getElementsByClassName('search-container');
  if (searchContainer.length > 0) {
    var form = document.getElementById("search-form");
    form.addEventListener("submit", handleFormSubmit);
    document.getElementsByClassName('search-container')[0].addEventListener('show', onShow);
  }

  function onShow() {
    makeSearchCall();
  }

  function handleFormSubmit(e) {
    e.preventDefault();
    makeSearchCall()
  }

  function makeSearchCall() {
    var searchResults = document.getElementsByClassName('search-results')[0];
    searchResults.innerHTML = SPINNER;

    var input = document.getElementById('search-term');

    var xhr = new XMLHttpRequest();

    xhr.addEventListener("load", handleCreateSuccess);
    xhr.addEventListener("error", handleCreateError);

    xhr.open("GET", "/_api/v1/search?s=" + encodeURIComponent(input.value));
    xhr.setRequestHeader("Accepts", "application/json");

    xhr.send();
  }

  function handleCreateSuccess(event) {
    var searchResults = document.getElementsByClassName('search-results')[0];
    searchResults.innerHTML = '';

    if (event.target.status === 200) {
      var results = JSON.parse(event.target.response.trim());
      if (!results || results.length === 0) {
        searchResults.innerHTML = NO_RESULTS;
        return;
      }
      var resultNodes = results.map(function(result) {
        var link = result['Link'];
        var url = result['URL'];
        return '<li>' + link + ': ' + '<a href="' + url + '">' + url + '</a>' + '</li>';
      });
      searchResults.innerHTML = '<ul>' + resultNodes.join('') + '</ul>';
    } else {
      handleCreateError(event);
    }
  }

  function handleCreateError(event) {
    console.log("error", event.target.status, event.target.response);
    // @TODO(jengler) 2016-2-22: Don't expose raw response to user.
    alert("Error: " + event.target.response);
  }
})(window, document)
