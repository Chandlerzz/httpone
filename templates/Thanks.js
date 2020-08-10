import React, { Component } from "react";
import { Link } from "react-router-dom";

export class Thanks extends Component {
  constructor(props) {
        super(props);
}
    render() {
        return <div class="container py-5">
    <div class="row text-center text-white mb-5">
        <div class="col-lg-7 mx-auto">
            <h1 class="display-4">Product List</h1>
        </div>
    </div>
    <div class="row">
        <div class="col-lg-8 mx-auto">

            <ul class="list-group shadow">

                <li class="list-group-item">

                    <div class="media align-items-lg-center flex-column flex-lg-row p-3">
                        <div class="media-body order-2 order-lg-1">
                            <h5 class="mt-0 font-weight-bold mb-2">OnePlus 7 Pro (Almond, 256 GB)</h5>
                            <p class="font-italic text-muted mb-0 small">Rear Camera|48MP (Primary)+ 8MP (Tele-photo)+16MP (ultrawide)| Front Camera|16 MP POP-UP Camera|8GB RAM|Android pie</p>
                            
                        </div>
                        <img src="https://i.imgur.com/6IUbEME.jpg" alt="Generic placeholder image" width="200" class="ml-lg-5 order-1 order-lg-2"></img>
                    </div>
                </li>
            </ul>
        </div>
    </div>
</div>
    }
}
