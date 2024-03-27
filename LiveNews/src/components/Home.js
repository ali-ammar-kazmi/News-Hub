import { React, Component } from "react";

export default class Home extends Component {
  render() {
  return (
    <>
      <div class="containter-fluid main_header">
        <div class="row">
          <div class="col-md-6 col-12 main_header_right align-items-center">
            <figure>
              <img
                src="static/about10.png"
                alt="student"
                className="img-fluid"
                title="TechStudent"
              />
            </figure>
          </div>
          <div class="col-md-6 col-12 main_header_left">
            <h1>Welcome to API-Hub</h1>
            <h2>
              Job Profile : <span class="text_clr"> Automation Engineer </span>
            </h2>
            <a href="https://www.linkedin.com/in/ali-ammar-kazmi-598701215" rel="noreferrer" target="_blank">
              <button>Reach Me</button>
            </a>
          </div>
        </div>
      </div>
    </>
  );
}
};
