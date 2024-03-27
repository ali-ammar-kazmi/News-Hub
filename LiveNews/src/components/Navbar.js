import { React, Component } from "react";
import { Link } from "react-router-dom";
import { categories } from "./constants";

export default class Navbar extends Component {
  render() {
    return (
      <>
        <nav className="navbar navbar-expand-lg bg-dark navbar-dark">
          <div className="container-fluid">
            <div className="navbar-brand display-6">{this.props.title}</div>
            <div
              className="collapse navbar-collapse"
              id="navbarSupportedContent"
            >
              <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                <li className="nav-item">
                  <Link className="nav-link active" aria-current="page" to="/Home">
                    {this.props.home}
                  </Link>
                </li>
                {categories.map((category) => (
                  <li className="nav-item" key={category}>
                    <Link className="nav-link" to={`/${category}`}>
                      {category}
                    </Link>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </nav>
      </>
    );
  }
}
