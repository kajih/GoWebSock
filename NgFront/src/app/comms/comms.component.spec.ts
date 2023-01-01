import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CommsComponent } from './comms.component';

describe('CommsComponent', () => {
  let component: CommsComponent;
  let fixture: ComponentFixture<CommsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CommsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CommsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
