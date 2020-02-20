import React, {Component} from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";
import * as Paths from "../../constants/routes"

import MemoryViewer from "../MemoryViewer"
import Game from "../Game"
import Console from "../Console"

export class App extends Component{
 
  componentWillMount(){
   
  }

  game(){
    return(<span>Game function</span>)
  }

  render(){
    return(
      <div>
        <Router>
          <Switch>
            <Route exact path={Paths.HOME}> <Game/> </Route>
            <Route path={Paths.MEMORY_VIEWER}> <MemoryViewer/> </Route>
            <Route path={Paths.CONSOLE}><Console/></Route>
          </Switch>
        </Router>

      </div>
      )
  }
}