import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MySpreadComponent } from './my-spread.component';

describe('MySpreadComponent', () => {
  let component: MySpreadComponent;
  let fixture: ComponentFixture<MySpreadComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MySpreadComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MySpreadComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
