(function(window, document, undefined) {
  "use strict";

  var SPINNER = '<i class="fa fa-2x fa-pulse fa-spinner"></i>';
  var NO_RESULTS = '<span>No results found</span>';

  window.onload = function() {
    var buttons = Array.from(document.getElementById('days-nav').getElementsByTagName('button'));
    buttons.forEach(function(button) {
      button.addEventListener('click', function() {
        buttons.forEach(function(btn) { btn.classList.remove('selected') });
        button.classList.add('selected');
        makeTopNCall(10, button.attributes.getNamedItem('data-days').value);
      });
    });
    makeTopNCall(10, 1);
  };

  function makeTopNCall(numResults, days) {
    var topNResults = document.getElementsByClassName('top-n-results')[0];
    topNResults.innerHTML = SPINNER;

    var xhr = new XMLHttpRequest();

    xhr.addEventListener("load", handleCreateSuccess);
    xhr.addEventListener("error", handleCreateError);

    xhr.open("GET", "/_api/v1/top_n?n=" + encodeURIComponent(numResults) + "&days=" + encodeURIComponent(days));
    xhr.setRequestHeader("Accepts", "application/json");

    xhr.send();
  }

  function handleCreateSuccess(event) {
    var topNResults = document.getElementsByClassName('top-n-results')[0];
    topNResults.innerHTML = '';

    if (event.target.status === 200) {
      var results = JSON.parse(event.target.response.trim());
      if (!results || results.length === 0) {
        topNResults.innerHTML = NO_RESULTS;
        return;
      }
      var resultNodes = results.map(function(result) {
        var link = result['Link'];
        var count = result['HitCount'];
        return '<tr><td><a href="/' + link + '">' + link + '</a></td><td>' + count + '</td></tr>';
      });
      var table = '<table><thead><tr><th>Link</th><th>Count</th></tr></thead><tbody>' + resultNodes.join('') + '</tbody></table>';
      topNResults.innerHTML = table;
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
