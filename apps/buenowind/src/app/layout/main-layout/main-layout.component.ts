import { ChangeDetectionStrategy, Component } from "@angular/core";
import { RouterLink, RouterOutlet } from "@angular/router";

@Component({
    selector: "bw-main-layout",
    imports: [RouterOutlet, RouterLink],
    template: `
        <div class="flex h-screen">
            <!-- Sidebar -->
            <nav class="w-64 bg-gray-800 p-4 text-white">
                <h2 class="text-xl font-bold">Brand</h2>
                <ul class="mt-4">
                    <li class="rounded px-3 py-2 hover:bg-gray-700">
                        <a routerLink="/dashboard">Dashboard</a>
                    </li>
                    <li class="rounded px-3 py-2 hover:bg-gray-700">
                        <a routerLink="/settings">Settings</a>
                    </li>
                </ul>
            </nav>

            <!-- Main Content -->
            <div class="flex flex-1 flex-col">
                <!-- Top Navigation -->
                <header class="flex items-center justify-between bg-white p-4 shadow">
                    <h1 class="text-lg font-semibold">Page Title</h1>
                    <button class="rounded bg-blue-500 px-4 py-2 text-white">Profile</button>
                </header>

                <!-- Page Content -->
                <main class="flex-1 p-4">
                    <router-outlet />
                </main>
            </div>
        </div>
    `,
    styles: `
        :host {
            display: block;
            height: 100%;
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class MainLayoutComponent {}
