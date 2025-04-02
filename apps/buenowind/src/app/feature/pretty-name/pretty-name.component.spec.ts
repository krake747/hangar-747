import { ComponentFixture, TestBed } from "@angular/core/testing";

import { PrettyNameComponent } from "./pretty-name.component";

describe("PrettyNameComponent", () => {
    let component: PrettyNameComponent;
    let fixture: ComponentFixture<PrettyNameComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [PrettyNameComponent]
        }).compileComponents();

        fixture = TestBed.createComponent(PrettyNameComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it("should create", () => {
        expect(component).toBeTruthy();
    });
});
