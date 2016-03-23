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
    var tenantUsers, userLocationGroups = [];
    var users = $.get('/dashboard/' + Cookies.get('tenantid') + '/users', function (result) {
        tenantUsers = result
    });

    var allowedGroups = $.get('/dashboard/' + Cookies.get('tenantid') + '/users/' + Cookies.get('username'), function (result) {
        userLocationGroups = result.apgroups
    });
    $.when(users, allowedGroups).done(function () {
        $.get('components/dashboard-usertable.html', function (template) {
                var rendered = Mustache.render(template, {
                    data: JSON.stringify(tenantUsers),
                    tenantDomain: Cookies.get('tenantdomain'),
                    roles: [{name: "Admin"}, {name: "DevOp"}, {name: "Dashboard User"}],
                    groups: userLocationGroups
                });
                $('#content-main').html(rendered);
            }
        );
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



