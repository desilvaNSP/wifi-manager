/**
 * Created by anuruddha on 9/13/15.
 */

function renderSidebar(username){
    $.get('components/sidebar.html', function(template) {
            $.get('/dashboard/users/1/' + username , function(result){
                var rendered = Mustache.render(template, {data: result});
                $('#side-navigation').html(rendered);
            });
        }
    );
}

function renderWifiUserTable(){
    $.get('components/wifi-usertable.html', function(template) {
            var users;
            $.get('/wifi/users', function(result){
                var rendered = Mustache.render(template, {data: result});
                $('#content-main').html(rendered);
            })
        }
    );
}

function renderDashboardUserTable(){
    $.get('components/dashboard-usertable.html', function(template) {
            var users;
            $.get('/dashboard/1/users', function(result){
                var rendered = Mustache.render(template, {data: result});
                $('#content-main').html(rendered);
            })
        }
    );
}

function renderDashBoard() {

    $.get('components/dashboard.html', function (template) {
        $.get('/wifi/locations',function(result){
            if (result){
                window.wifilocation = result[0].locationid;
                window.wifilocationlist = result
            }else{
                window.wifilocation = "default";
            }
            var rendered = Mustache.render(template, {locations:result});
            $('#content-main').html(rendered)
        });
    })
}


function renderLocations() {
    $.get('components/locations.html', function (template) {
        $.get('/wifi/locations',function(result){
            var rendered = Mustache.render(template, {data: JSON.stringify(result)});
            $('#content-main').html(rendered)
        });
    })
}


