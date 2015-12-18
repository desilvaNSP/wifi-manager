function renderSidebar(username){
    $.get('components/sidebar.html', function(template) {
            $.get('/dashboard/users/'+Cookies.get('tenantId')+ '/' + username , function(result){
                var rendered = Mustache.render(template, {data: result});
                $('#side-navigation').html(rendered);
            });
        }
    );
}

function renderWifiUserTable(){
    var loading = $.get('components/loading.html', function(template) {
            var rendered = Mustache.render(template);
            $('#content-main').html(rendered);
        }
    );

    $.when(loading).done(function(){
        $.get('components/wifi-usertable.html', function(template) {
                var users;
                $.get('/wifi/users', function(result){
                    var rendered = Mustache.render(template, {data: result});
                    $('#content-main').html(rendered);
                })
            }
        );
    });
}

function renderDashboardUserTable(){
    var tenantRoles, tenantUsers;
    var roles = $.get('/dashboard/'+ Cookies.get('tenantId') +'/roles', function(result){
        tenantRoles = result;
    });
    var users = $.get('/dashboard/'+ Cookies.get('tenantId') +'/users', function(result){
        tenantUsers = result
    });

    $.when(roles,users).done(function(){
        $.get('components/dashboard-usertable.html', function(template) {
                var rendered = Mustache.render(template, {data: tenantUsers, roles : JSON.parse(tenantRoles)});
                $('#content-main').html(rendered);
            }
        );
    });
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

function renderDashboardList() {
    $.get('components/dashboard-list.html', function (template) {

            var rendered = Mustache.render(template, {});
            $('#content-main').html(rendered)
    })
}


