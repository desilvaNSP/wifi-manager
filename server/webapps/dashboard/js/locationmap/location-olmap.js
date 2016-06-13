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

function countryVectorLayer(map){
    var vectorCountriesSource = new ol.source.Vector({
        url: 'map-data/geojson/countries.geojson',
        format: new ol.format.GeoJSON()
    });
    var vectorCountriesLayer = new ol.layer.Vector({
        source: vectorCountriesSource,
    });

    map.addLayer(vectorCountriesLayer);

}

function enableGeoJsonDragDrop(){
    var defaultStyle = {
        'Point': new ol.style.Style({
            image: new ol.style.Circle({
                fill: new ol.style.Fill({
                    color: 'rgba(255,255,0,0.5)'
                }),
                radius: 5,
                stroke: new ol.style.Stroke({
                    color: '#ff0',
                    width: 1
                })
            })
        }),
        'LineString': new ol.style.Style({
            stroke: new ol.style.Stroke({
                color: '#f00',
                width: 3
            })
        }),
        'Polygon': new ol.style.Style({
            fill: new ol.style.Fill({
                color: 'rgba(0,255,255,0.5)'
            }),
            stroke: new ol.style.Stroke({
                color: '#0ff',
                width: 1
            })
        }),
        'MultiPoint': new ol.style.Style({
            image: new ol.style.Circle({
                fill: new ol.style.Fill({
                    color: 'rgba(255,0,255,0.5)'
                }),
                radius: 5,
                stroke: new ol.style.Stroke({
                    color: '#f0f',
                    width: 1
                })
            })
        }),
        'MultiLineString': new ol.style.Style({
            stroke: new ol.style.Stroke({
                color: '#0f0',
                width: 3
            })
        }),
        'MultiPolygon': new ol.style.Style({
            fill: new ol.style.Fill({
                color: 'rgba(0,0,255,0.5)'
            }),
            stroke: new ol.style.Stroke({
                color: '#00f',
                width: 1
            })
        })
    };

    var styleFunction = function(feature, resolution) {
        var featureStyleFunction = feature.getStyleFunction();
        if (featureStyleFunction) {
            return featureStyleFunction.call(feature, resolution);
        } else {
            return defaultStyle[feature.getGeometry().getType()];
        }
    };

    var dragAndDropInteraction = new ol.interaction.DragAndDrop({
        formatConstructors: [
            ol.format.GPX,
            ol.format.GeoJSON,
            ol.format.IGC,
            ol.format.KML,
            ol.format.TopoJSON
        ]
    });

    var map = new ol.Map({
        interactions: ol.interaction.defaults().extend([dragAndDropInteraction]),
        layers: [
            new ol.layer.Tile({
                source: new ol.source.BingMaps({
                    imagerySet: 'Aerial',
                    key: 'Your Bing Maps Key from http://www.bingmapsportal.com/ here'
                })
            })
        ],
        target: 'world-map',
        view: new ol.View({
            center: [0, 0],
            zoom: 2
        })
    });

    dragAndDropInteraction.on('addfeatures', function(event) {
        var vectorSource = new ol.source.Vector({
            features: event.features
        });
        map.addLayer(new ol.layer.Vector({
            source: vectorSource,
            style: styleFunction
        }));
        map.getView().fit(
            vectorSource.getExtent(), /** @type {ol.Size} */ (map.getSize()));
    });

    var displayFeatureInfo = function(pixel) {
        var features = [];
        map.forEachFeatureAtPixel(pixel, function(feature) {
            features.push(feature);
            console.log(feature)
        });
        if (features.length > 0) {
            var info = [];
            var i, ii;
            for (i = 0, ii = features.length; i < ii; ++i) {
                info.push(features[i].get('STATE_NAME'));
            }
            document.getElementById('info').innerHTML = info.join(', ') || '&nbsp';
        } else {
            document.getElementById('info').innerHTML = '&nbsp;';
        }
    };

    map.on('pointermove', function(evt) {
        if (evt.dragging) {
            return;
        }
        var pixel = map.getEventPixel(evt.originalEvent);
        displayFeatureInfo(pixel);
    });

    map.on('click', function(evt) {
        displayFeatureInfo(evt.pixel);
    });
}