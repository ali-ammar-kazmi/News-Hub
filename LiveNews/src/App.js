import { React, Component } from "react";
import Navbar from "./components/Navbar";
import News from "./components/News";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { categories } from "./components/constants";
import Home from "./components/Home";

export default class App extends Component {
  renderRoutes = () => {
    return categories.map((category) => {
      return (
        <Route
          key={category}
          exact
          path={`/${category}`}
          element={
            category === "Home" ? (
              <Home />
            ) : (
              <News key={category} category={category} />
            )
          }
        />
      );
    });
  };

  render() {
    return (
      <>
        <Router>
          <Navbar />
          <Routes>{this.renderRoutes()}</Routes>
        </Router>
      </>
    );
  }
}
