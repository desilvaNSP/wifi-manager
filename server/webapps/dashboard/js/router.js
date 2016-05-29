function renderSidebar(username) {
    $.get('components/sidebar.html', function (template) {
            $.get('/dashboard/' + Cookies.get('tenantid') + '/users/' + username, function (result) {
                var rendered = Mustache.render(template, {data: result});
                $('#side-navigation').html(rendered);
            });
        }
    );
}

function renderWifiUserTable() {
    var loading = $.get('components/loading.html', function (template) {
            var rendered = Mustache.render(template);
            $('#content-main').html(rendered);
        }
    );

    $.when(loading).done(function () {
        $.get('components/wifi-usertable.html', function (template) {
                var rendered = Mustache.render(template, {});
                $('#content-main').html(rendered);
            }
        );
    });
}

function renderDashboardUserTable() {
    var editUserFromValidator;
    var tenantUsers, userLocationGroups = []
    var userPermissions = [];
    var users = $.get('/dashboard/' + Cookies.get('tenantid') + '/users', function (result) {
        tenantUsers = result
    });

    var allowedGroups = $.get('/dashboard/' + Cookies.get('tenantid') + '/users/' + Cookies.get('username'), function (result) {
        userLocationGroups = result.apgroups
    });

    var allowedUserScopes = $.get('/dashboard/' + Cookies.get('tenantid') + '/permissions', function (data) {
        for(var i in data ){
            var obj = {};
            var scopename = i.replace('_'," ")
            obj.id = i;
            obj.name = scopename.charAt(0).toUpperCase() + scopename.slice(1);
            obj.actions = data[i];
            userPermissions.push(obj);
        }
    });


    $.when(users, allowedGroups, allowedUserScopes).done(function () {
       var renderUsertable = $.get('components/dashboard-usertable.html', function (template) {
                var rendered = Mustache.render(template, {
                    data: JSON.stringify(tenantUsers),
                    tenantDomain: Cookies.get('tenantdomain'),
                    roles: [{name: "Admin"}, {name: "DevOp"}, {name: "Dashboard User"}],
                    groups: userLocationGroups,
                    userscopes:userPermissions
                });
                $('#content-main').html(rendered);
            }
        );

        $.when(renderUsertable).done(function(){
            $.get('components/dashboard-userupdate-modal.html', function (template) {
                var rendered = Mustache.render(template, {
                    groups:userLocationGroups,
                    userscopes:userPermissions,
                    tenantDomain:Cookies.get('tenantdomain')
                });
                $('#userupdate-modal-content').html(rendered);
            });
        });
    });
}

function renderLocations() {
    $.get('components/locations.html', function (template) {
        var rendered = Mustache.render(template, {});
        $('#content-main').html(rendered)
    })
}

function renderDashboardList() {
    $.get('components/dashboard-list.html', function (template) {
        var rendered = Mustache.render(template, {});
        $('#content-main').html(rendered)
    })
}

function renderEmptyPage() {
    $.get('components/emptypage.html', function (template) {
        var rendered = Mustache.render(template, {});
        $('#content-main').html(rendered)
    })
}
function renderReportPage() {
    $.get('components/emptypage.html', function (template) {
        var rendered = Mustache.render(template, {});
        $('#content-main').html(rendered)
    })
}

function renderProfilePage() {
    $.get('components/profile.html', function (template) {
        var rendered = Mustache.render(template, {});
        $('#content-main').html(rendered)
    })
}

function renderAAAPage() {
    $.get('components/radius-aaamonitor.html', function (template) {
        var rendered = Mustache.render(template, {});
        $('#content-main').html(rendered)
    })
}

function renderAlertDefinations() {
    $.get('components/emptypage.html', function (template) {
        var rendered = Mustache.render(template, {});
        $('#content-main').html(rendered)
    })
}



