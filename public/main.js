// create an array with nodes





$(document).ready(function (){

    var json;
    //var dataset = 'data/breeze-breeze-token.json';
    var dataset = 'data/pacnem-cheese.json';
    //var dataset = 'data/banco-coin.json';
    // var dataset = 'data/dim-coin.json';
    // var dataset = 'data/gold-gold.json';
    //var dataset = 'data/nemventory.product-beginners_fishing_rod.json';

    $.getJSON(dataset, function (response) {
        json = response;
        console.log(json);
        $(window).trigger('JSONready');
    });



    // $(window).on('JSONready', function () {
    //
    //     console.log("Json ready")
    //
    //    var g = {
    //        nodes: [],
    //        edges: []
    //    };
    //
    //
    //
    //     var color1 = '#ff2700'
    //     var color2 = '#5b64ff'
    //     var color3 = '#2dff07'
    //     var color4 = '#ff00e4'
    //     var color5 = '#000000'
    //     var color6 = '#ff9400'
    //
    //
    //     var i;
    //     for ( i = 0; i < json.nodes.length; i++ ){
    //         n = json.nodes[i];
    //         g.nodes.push({
    //             // Main attributes:
    //             id: n.id,
    //             label: n.label,
    //             // Display attributes:
    //             x: Math.cos(Math.PI * 2 * i / json.nodes.length),
    //             y: Math.sin(Math.PI * 2 * i / json.nodes.length),
    //             size: n.value,
    //             color: eval("color"+n.group)
    //         })
    //     }
    //
    //     for ( i = 0; i < json.edges.length; i++ ){
    //         e = json.edges[i];
    //         g.edges.push({
    //             id: e.from+e.to+i,
    //             // Reference extremities:
    //             source: e.from,
    //             target: e.to,
    //             label: String(e.value),
    //             color: eval("color"+e.group),
    //             type: 'curvedArrow'
    //         });
    //     }
    //
    //     // Instantiate sigma:
    //     s = new sigma({
    //         graph: g,
    //         renderer: {
    //             container: document.getElementById('network'),
    //             type: 'canvas'
    //         },
    //         settings: {
    //
    //         }
    //     });
    //     // s.startForceAtlas2();
    //     //
    //     // setTimeout(function () {
    //     //     s.stopForceAtlas2();
    //     // }, 3000)(s)
    //
    // });

    $(window).on('old2', function () {

        console.log("Json ready")

        var graph = Viva.Graph.graph();

        var i;
        for ( i = 0; i < json.edges.length; i++ ){
            console.log(i);
            graph.addLink(json.edges[i].to, json.edges[i].from);
        }

        var graphics = Viva.Graph.View.cssGraphics();
        graphics.node(function(node) {
            // The function is called every time renderer needs a ui to display node
            return Viva.Graph.css('')
                .attr('color', 24)
                .attr('height', 24)

        })
            .placeNode(function(nodeUI, pos){
                // Shift image to let links go to the center:
                nodeUI.attr('x', pos.x - 12).attr('y', pos.y - 12);
            });

        var renderer = Viva.Graph.View.renderer(graph, {
            container: document.getElementById('network'),
            graphics : graphics
        });
        renderer.run();




    });




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
            configure: {
                container: document.getElementById('config'),
            },
            nodes: {
            },
            layout:{
                hierarchical: false
            },
        };
        var network = new vis.Network(
            document.getElementById('network'),
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