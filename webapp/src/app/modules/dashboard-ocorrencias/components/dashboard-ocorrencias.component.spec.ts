import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DashboardOcorrenciasComponent } from './dashboard-ocorrencias.component';

describe('DashboardOcorrenciasComponent', () => {
  let component: DashboardOcorrenciasComponent;
  let fixture: ComponentFixture<DashboardOcorrenciasComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DashboardOcorrenciasComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DashboardOcorrenciasComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
