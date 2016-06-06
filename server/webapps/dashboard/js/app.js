$(document).ajaxError(function (event, jqxhr, settings, thrownError) {
    toastr.options = {
        "closeButton": true,
        "debug": false,
        "progressBar": true,
        "preventDuplicates": true,
        "positionClass": "toast-top-right",
        "onclick": null,
        "showDuration": "400",
        "hideDuration": "1500",
        "timeOut": "1500",
        "extendedTimeOut": "1000",
        "showEasing": "swing",
        "hideEasing": "linear",
        "showMethod": "fadeIn",
        "hideMethod": "fadeOut"
    };
    if (thrownError == 'Unauthorized') {
        toastr.error("Please check your username and password")
        toastr.options.onHidden = function () {
            window.location.href = "/dashboard/login"
        }
    }
    if (thrownError == 'Forbidden') {
        toastr.options.onHidden = function () {
        };
        toastr.error("You are not Authorized to perform this operation")
    }
});

$(document).ajaxSend(function (event, request, settings) {
    request.setRequestHeader("Authorization", "Bearer " + Cookies.get("jwt"));
});

$(document).ready(function () {
    if (!Cookies.get("jwt")) {
        //toastr.error("Unauthorised please login")
        window.location.href = "/dashboard/login"
        return
    } else {
        storeLoggedInUserPermissions()
    }
    renderSidebar(Cookies.get("username"))
    renderDashboardList();

    // Add body-small class if window less than 768px
    if ($(this).width() < 769) {
        $('body').addClass('body-small')
    } else {
        $('body').removeClass('body-small')
    }

    // Close menu in canvas mode
    $('.close-canvas-menu').click(function () {
        $("body").toggleClass("mini-navbar");
        SmoothlyMenu();
    });

    // Initialize slimscroll for small chat
    $('.small-chat-box .content').slimScroll({
        height: '234px',
        railOpacity: 0.4
    });

    // Minimalize menu
    $('.navbar-minimalize').click(function () {
        $("body").toggleClass("mini-navbar");
        SmoothlyMenu();
    });

    // Tooltips demo
    $('.tooltip-demo').tooltip({
        selector: "[data-toggle=tooltip]",
        container: "body"
    });

    // Move modal to body
    // Fix Bootstrap backdrop issu with animation.css
    $('.modal').appendTo("body");

    // Full height of sidebar
    function fix_height() {
        var heightWithoutNavbar = $("body > #wrapper").height() - 61;
        $(".sidebard-panel").css("min-height", heightWithoutNavbar + "px");

        var navbarHeigh = $('nav.navbar-default').height();
        var wrapperHeigh = $('#page-wrapper').height();

        if (navbarHeigh > wrapperHeigh) {
            $('#page-wrapper').css("min-height", navbarHeigh + "px");
        }

        if (navbarHeigh < wrapperHeigh) {
            $('#page-wrapper').css("min-height", $(window).height() + "px");
        }

    }

    fix_height();

    // Fixed Sidebar
    $(window).bind("load", function () {
        if ($("body").hasClass('fixed-sidebar')) {
            $('.sidebar-collapse').slimScroll({
                height: '100%',
                railOpacity: 0.9
            });
        }
    })

    // Move right sidebar top after scroll
    $(window).scroll(function () {
        if ($(window).scrollTop() > 0 && !$('body').hasClass('fixed-nav')) {
            $('#right-sidebar').addClass('sidebar-top');
        } else {
            $('#right-sidebar').removeClass('sidebar-top');
        }
    });

    $(document).bind("load resize scroll", function () {
        if (!$("body").hasClass('body-small')) {
            fix_height();
        }
    });

    $("[data-toggle=popover]")
        .popover();

    // Add slimscroll to element
    $('.full-height-scroll').slimscroll({
        height: '100%'
    })

    $(".logout").on('click', function (event) {
        event.stopImmediatePropagation();
        $.ajax({
            url: '/dashboard/logout',
            type: 'POST',
            success: function () {
                Cookies.remove('jwt', { secure: true });
                window.location.href = "/dashboard/login"
            },
            error: function (e) {
            }
        });
    })
});


// Minimalize menu when screen is less than 768px
$(window).bind("resize", function () {
    if ($(this).width() < 769) {
        $('body').addClass('body-small')
    } else {
        $('body').removeClass('body-small')
    }
});

// Collapse ibox function
$(document.body).on('click', '.collapse-link', function () {
    var ibox = $(this).closest('div.ibox');
    var button = $(this).find('i');
    var content = ibox.find('div.ibox-content');
    content.slideToggle(200);
    button.toggleClass('fa-chevron-up').toggleClass('fa-chevron-down');
    ibox.toggleClass('').toggleClass('border-bottom');
    setTimeout(function () {
        ibox.resize();
        ibox.find('[id^=map-]').resize();
    }, 50);
});

// Close ibox function
$(document).on('click', '.close-link', function () {
    var content = $(this).closest('div.ibox');
    content.remove();
});

// For demo purpose - animation css script
function animationHover(element, animation) {
    element = $(element);
    element.hover(
        function () {
            element.addClass('animated ' + animation);
        },
        function () {
            //wait for animation to finish before removing classes
            window.setTimeout(function () {
                element.removeClass('animated ' + animation);
            }, 2000);
        });
}

function SmoothlyMenu() {
    if (!$('body').hasClass('mini-navbar') || $('body').hasClass('body-small')) {
        // Hide menu in order to smoothly turn on when maximize menu
        $('#side-menu').hide();
        // For smoothly turn on menu
        setTimeout(
            function () {
                $('#side-menu').fadeIn(500);
            }, 100);
    } else if ($('body').hasClass('fixed-sidebar')) {
        $('#side-menu').hide();
        setTimeout(
            function () {
                $('#side-menu').fadeIn(500);
            }, 300);
    } else {
        // Remove all inline style from jquery fadeIn function to reset menu state
        $('#side-menu').removeAttr('style');
    }
}

function storeLoggedInUserPermissions() {
    $.get('/dashboard/' + Cookies.get('tenantid') + '/users/' + Cookies.get('username'), function (result) {
        Cookies.set('userpermissions', result.permissions)
    })
}

function getPieChartData(data) {
    return pieChartData = $.map(data, function (obj, i) {
        return [{"name": obj.name, "y": obj.value}];
    });
}

function convertToHighChartSeries(arr, devider) {
    return data = $.map(arr, function (val, i) {
        return [[moment.utc(val.name, 'YYYY-M-D H:m:s').valueOf(), val.value / devider]];
    });
}

function renderTimeSeries(element, toolTipHeader, yAxisTitle, dataSeries) {
    $(element).highcharts({
        chart: {
            zoomType: 'x'
        },
        title: {
            text: ''
        },
        xAxis: {
            type: 'datetime',
            dateTimeLabelFormats: { // don't display the dummy year
                month: '%e. %b',
                year: '%b'
            },
            title: {
                text: 'Date'
            }
        },
        yAxis: {
            title: {
                text: yAxisTitle
            },
            min: 0
        },
        tooltip: {
            headerFormat: '<b>' + toolTipHeader + '</b><br>',
            pointFormat: '{point.x:%e. %b}: {point.y:.2f}'
        },
        plotOptions: {
            spline: {
                marker: {
                    enabled: true
                }
            }
        },
        exporting: {
            sourceWidth: 1600,
            sourceHeight: 500,
            chartOptions: {
                subtitle: null
            }
        },
        series: dataSeries,
        credits: {
            enabled: false
        }
    });
}


