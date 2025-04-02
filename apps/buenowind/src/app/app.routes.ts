import { Routes } from "@angular/router";
import { MainLayoutComponent } from "./layout/main-layout/main-layout.component";

export const routes: Routes = [
    // {
    //     path: "",
    //     pathMatch: "full", // Ensures exact match for empty path
    //     redirectTo: "home"
    // },
    {
        path: "",
        component: MainLayoutComponent,
        children: [
            {
                path: "home", // shorter import() because of use of "export default"
                loadChildren: () => import("./feature/home/home.routes")
            }
        ]
    },
    {
        path: "**", // fallback route (can be used to display dedicated 404 lazy feature)
        redirectTo: ""
    }
];
