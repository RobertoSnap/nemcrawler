// create an array with nodes





$(document).ready(function (){

    var JSON;
    $.getJSON('data/breeze_breeze-token.json', function (response) {
        JSON = response;
        console.log(JSON);
        $(window).trigger('JSONready');
    });


    console.log("Doc ready")

    $(window).on('JSONready', function () {

        console.log("Json ready")

        var nodes = new vis.DataSet(JSON.nodes);

        var edges = new vis.DataSet(JSON.edges);

        var data = {
            nodes: nodes,
            edges: edges
        };

        var options = {
            autoResize: true,
            height: '100%',
            width: '100%',
            nodes: {
                shape: "triangle",
            },
            layout:{
                hierarchical: false
            },
        };
        var network = new vis.Network(
            document.getElementById('mynetwork'),
            data,
            options
        );

        var test = {
            joinCondition:function(nodeOptions) {
                return nodeOptions.Address;
            }
        }
        network.clustering.cluster(test);


    });



})