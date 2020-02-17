import React, {Component} from "react"
import {TextField} from "@material-ui/core"
import Page from "../Page"
import PropTypes from 'prop-types';
import {List, AutoSizer} from "react-virtualized"
import {Legend} from './legend'
import * as LegendColors from '../../constants/memoryMapColours'
const handlers = require('../../ipcCalbacks')


const electron = window.require('electron');
const fs = electron.remote.require('fs');
const ipcRenderer  = electron.ipcRenderer;

export default class MemoryViewer extends Component{

    constructor(props){
        super(props);

        this.rowRenderer = this.rowRenderer.bind(this)
        this.onSearchKeyUp = this.onSearchKeyUp.bind(this);

        ipcRenderer.on('info' , function(event , data){ debugger; console.log(data.msg) });

        let mem = []
        for(let i = 0; i < 0xFFFF; i++){
            mem.push(0);
        }

        this.state = {
            content: mem,
            scrollToIndex: null
        };
    }

    pad(n, width, z) {
        z = z || '0';
        n = n + '';
        return n.length >= width ? n : new Array(width - n.length + 1).join(z) + n;
      }

    rowRenderer({ index, key, style }){
        let color = this.getBackgroundColor(index);

        let isIndexSelectedClass = index===this.state.scrollToIndex ? "bold large" : ''

        return (
            <div key={key} style={style}>
                <div className={"container " + isIndexSelectedClass} style={{backgroundColor:color, paddingBottom:2}}>
                    <span className={"container justify-center"} style={{flex: 2}}>
                            0x{this.pad(index.toString(16), 4).toUpperCase()}
                    </span>
                    <span className={"container justify-center"} style={{flex: 3}}>
                            0x{this.pad(this.state.content[index],4).toUpperCase()}
                    </span>
                </div>
                
            </div>            
          );
    }

    getBackgroundColor(i){
        if(i<= 0x3FFF)
            return LegendColors.CARTRIDGE_BANK00
        if(i <= 0x7FFF)
            return LegendColors.CARTRIDGE_BANKN
        if(i <= 0x9FFF)
            return LegendColors.VRAM
        if(i <= 0xBFFF)
            return LegendColors.EXTERNAL_RAM
        if(i <= 0xDFFF)
            return LegendColors.RAM
        if(i <= 0xFDFF)
            return LegendColors.MIRRORS
        if(i <= 0xFE9F)
            return LegendColors.OAM
        if(i <= 0xFEFF)
            return LegendColors.UNUSABLE
        if(i <= 0xFF7F)
            return LegendColors.IO
        if(i <= 0xFFFE)
            return LegendColors.HRAM
        if(i <= 0xFFFF)
            return LegendColors.INTERRUPTS
    }

    renderLegend(){
        return(
            <div className={"container-column"} style={{marginTop: 10, marginBottom: 10}}>
                <div className={"container"} style={{justifyContent: 'space-between'}}>
                    <Legend name='Cartridge' color={LegendColors.CARTRIDGE_BANK00}/>
                    <Legend name='External RAM' color={LegendColors.EXTERNAL_RAM}/>
                    <Legend name='VRAM' color={LegendColors.VRAM}/>
                </div>
                <div className={"container"} style={{justifyContent: 'space-between'}}>
                    <Legend name='Mirrors' color={LegendColors.MIRRORS}/>
                    <Legend name='RAM' color={LegendColors.RAM}/>
                    <Legend name='OAM' color={LegendColors.OAM}/>
                </div>
                <div className={"container"} style={{justifyContent: 'space-between'}}>
                    <Legend name='Unusable' color={LegendColors.UNUSABLE}/>
                    <Legend name='IO' color={LegendColors.IO}/>
                    <Legend name='HRAM' color={LegendColors.HRAM}/>
                </div>
            </div>
        )
    }

    onSearchKeyUp(e){
        if(e.target.value === ''){
            this.setState({
                scrollToIndex: null
            })
        }

        let val = parseInt(e.target.value);

        //See if we're searching by hex value
        if(!isNaN(val)){
            this.setState({
                scrollToIndex: val
            })
        }
    }

    onMemoryUpdated(data){
        this.setState({
            content: [...this.state.content, data[0], data[1]]
        })
    }

    render(){
        return(
            <Page style={{marginLeft: 10, marginRight: 10}}>
                <span className={"center title"}>MEMORY  VIEWER</span>
                <TextField id="outlined-basic" size={"small"} onKeyUp={this.onSearchKeyUp} label="Search by address" variant="outlined" />

                {/** Header */}
                <div className={"container"} style={{marginTop: 10, marginBottom:5}}>
                    <span className={"container justify-center super"} style={{flex: 2}}>Address</span>
                    <span className={"container justify-center super"} style={{flex: 3}}>Content</span>
                </div>

                <div style={{ display: 'flex', flexGrow:1, border:'1px black solid'}}>  
                    <div style={{ flex: '1 1 auto' }}>
                        <AutoSizer>
                            {({height, width}) => (
                                <List
                                    height={height}
                                    rowCount={this.state.content.length}
                                    rowHeight={20}
                                    rowRenderer={this.rowRenderer}
                                    width={width}
                                    scrollToIndex={this.state.scrollToIndex}
                                />
                            )}
                        </AutoSizer>                
                    </div>
                </div>

                {this.renderLegend()}

                
                
            </Page>
            )
    }
}

MemoryViewer.propTypes = {
    content: PropTypes.array,
    scrollToIndex: PropTypes.number
}