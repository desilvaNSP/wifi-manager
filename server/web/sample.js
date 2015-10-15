/**
 * Created by anuruddha on 9/14/15.
 */

function loadUser() {
    $.get('templates/sample.mst', function(template) {
        var rendered = Mustache.render(template, {name: "Luke"});
        $('#target').html(rendered);
    });
}
$(document).ready(function () {

    loadUser();
});
