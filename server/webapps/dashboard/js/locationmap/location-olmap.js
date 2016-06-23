var markers = [];


function updateMarkers(data, map, iconStyle){
    markers = $.map(data, function(item){
        return new ol.Feature({
            geometry: new ol.geom.Point(ol.proj.transform([item.latitude,item.longitude], 'EPSG:4326',
                'EPSG:3857')),
            apname: item.apname,
            bssid: item.bssid,
            mac: item.mac,
            groupname:item.groupname,
            ssid:item.ssid
        });
    });

    var vectorSource = new ol.source.Vector({
        features: markers //add an array of features
    });

    var vectorLayer = new ol.layer.Vector({
        source: vectorSource,
        style: iconStyle
    });

    map.addLayer(vectorLayer);
    initializePopoverOnMarker(map)
}

function initializePopoverOnMarker(map){
    var element = document.getElementById('popup');

    var popup = new ol.Overlay({
        element: element,
        positioning: 'bottom-center',
        stopEvent: false
    });
    map.addOverlay(popup);
    // display popup on click
    map.on('click', function(evt) {
        var feature = map.forEachFeatureAtPixel(evt.pixel,
            function(feature) {
                return feature;
            });
        if (feature) {
            var renderdContent = "";
            $.get('components/widgets/location-mapmarker.html', function (template) {
                renderdContent = Mustache.render(template, feature.U);
                popup.setPosition(evt.coordinate);
                $(element).popover({
                    'placement': 'center',
                    'html': true,
                    'content': renderdContent
                });
                $(element).popover('show');
            });
        } else {
            $(element).popover('destroy');
        }
    });

    var cursorHoverStyle = "pointer";
    var target = map.getTarget();

    var jTarget = typeof target === "string" ? $("#"+target) : $(target);

    map.on("pointermove", function (event) {
        var mouseCoordInMapPixels = [event.originalEvent.offsetX, event.originalEvent.offsetY];
        var hit = map.forEachFeatureAtPixel(mouseCoordInMapPixels, function (feature, layer) {
            return true;
        });
        if (hit) {
            jTarget.css("cursor", cursorHoverStyle);
            $(element).popover('destroy');
        } else {
            jTarget.css("cursor", "");
        }
    });
}