import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { OcorrenciasComponent } from './components';
import { LayoutModule } from '../layout';

@NgModule({
  declarations: [
    OcorrenciasComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    RouterModule,
    LayoutModule
  ]
})
export class OcorrenciasModule { }
