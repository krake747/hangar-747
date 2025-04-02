import { ComponentFixture, TestBed } from "@angular/core/testing";

import { JohnComponent } from "./john.component";

describe("JohnComponent", () => {
    let component: JohnComponent;
    let fixture: ComponentFixture<JohnComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [JohnComponent]
        }).compileComponents();

        fixture = TestBed.createComponent(JohnComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it("should create", () => {
        expect(component).toBeTruthy();
    });
});
