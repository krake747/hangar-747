import { Routes } from "@angular/router";
import { HomeComponent } from "./home.component";

export default [
    {
        path: "",
        providers: [],
        component: HomeComponent,
        children: [
            {
                path: "tom",
                loadChildren: () => import("../john/john.routes")
            }
        ]
    },
    {
        path: "john",
        loadChildren: () => import("../john/john.routes")
    },
    {
        path: "p/:name",
        loadChildren: () => import("../pretty-name/pretty-name.routes")
    }
] as Routes;
