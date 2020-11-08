let pannel = {
    _divID: undefined,
    _width: undefined,
    _height: undefined,
    _margin: {left: 10, right: 10, top: 10, bottom: 10},
    _svg: undefined,
    left : 0,
    squarewidth: undefined,
    size: 20,
    initialize: function (divID) {
        self = this;
        self._divID = divID;
        self._width = $('#' + divID ).width()/3 - self._margin.left;
        self._height = $('#' + divID ).height() - self._margin.top;
        self.left = parseInt($('#' + divID ).width()/3)+ self._margin.left;

        let svg = d3.select('#' + divID )
            .append('svg')
            .attr('id', divID + '_svg')
            .attr('transform', 'translate(' + self.left+ ',' + self._margin.top + ')')
            .attr('width', self._width)
            .attr('height', self._height);

        self._svg = svg;
        self.squarewidth = parseInt(self._width/self.size);
        self.initPannel();
    },
    initPannel: function() {
        let i = 0;
        let j = 0;
        let svg = d3.select('#' + self._divID + '_svg');
        for (i=0;i<self.size;i++)
        { 
            for (j=0; j<self.size;j++) {
                let x = j*self.squarewidth;
                let y = self._margin.top + i*self.squarewidth;
                svg.append('g').attr('class', 'sub_square').append('rect')
                .attr('id','rect'+i+'_'+j)
                .attr('x', x)
                .attr('y', y)
                .attr('width', self.squarewidth)
                .attr('height', self.squarewidth)
                .style('stroke-width',1)
                .style('stroke','black')
                .attr('fill','none');
            }
        }
        
    },
    updatePannel: function(data) {
        //console.log(data);
        let fillMap = {0:'none',1:'gray',3:'black',10:'brown'}
        for (i=0;i<self.size;i++)
        { 
            for (j=0; j<self.size;j++) {
                let d = data[i][j];
                d3.select('#rect'+i+'_'+j).attr('fill',fillMap[d]);
            }
        }
    },
    reinitPannel: function() {
        //console.log(data);
        for (i=0;i<self.size;i++)
        { 
            for (j=0; j<self.size;j++) {
                d3.select('#rect'+i+'_'+j).attr('fill','none');
            }
        }
    }

}
