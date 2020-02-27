import React, {Component} from "react"
const { ipcRenderer } = window.require('electron');

export default class Console extends Component{

  constructor(props){
    super(props);

    ipcRenderer.on('log' , (event , data) => this.onLogReceived(data));

    this.state = {
        logs: [],
    };
  }

  onLogReceived = (data) => {
    this.state.logs.push(Buffer.from(data).toString())
    this.setState({
      logs : this.state.logs
    })
  }
 
  render(){
    return(
      <div style={{backgroundColor: "grey"}}>
      <p className={"white"}>{this.state.logs.join("\n")}</p>
      </div>
      )
  }
}