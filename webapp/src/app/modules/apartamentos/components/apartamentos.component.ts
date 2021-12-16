import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';

import { ApartamentosService } from '../services';
import { ApartamentoModel } from '../models';

@Component({
  selector: 'app-apartamentos',
  templateUrl: './apartamentos.component.html',
  styleUrls: ['./apartamentos.component.css']
})
export class ApartamentosComponent implements OnInit {

  apartamento!: ApartamentoModel;
  apartamentos: any = [];
  apt_selec: ApartamentoModel = {};
  apartamento_id!: number;

  constructor(
    private apartamentosService: ApartamentosService
  ) { }

  ngOnInit(): void {
    this.buscarTodosApartamentos();
  }

  apartamentoForm = new FormGroup({
    identificador: new FormControl(null, [
      Validators.required
    ]),
    descricao: new FormControl(null, [
      Validators.required
    ])
  });

  apartamentoModalForm = new FormGroup({
    identificador_modal: new FormControl(null, [
      Validators.required
    ]),
    descricao_modal: new FormControl(null, [
      Validators.required
    ])
  });

  get identificador(): any {
    return this.apartamentoForm.get('identificador');
  }

  get descricao(): any {
    return this.apartamentoForm.get('descricao');
  }

  get identificador_modal(): any {
    return this.apartamentoModalForm.get('identificador_modal');
  }

  get descricao_modal(): any {
    return this.apartamentoModalForm.get('descricao_modal');
  }

  cadastrarApartamento(): void {
    this.apartamento = { descricao: this.descricao.value, identificador: this.identificador.value };

    this.apartamentosService.cadastrarApartamento(this.apartamento)
      .subscribe(
        response => {
          this.buscarTodosApartamentos();
          this.apartamentoForm.reset();
        },
        error => { }
      );
  }

  buscarTodosApartamentos(): void {
    this.apartamentosService.buscarTodosApartamentos()
      .subscribe(
        response => {
          this.apartamentos = [];
          response.forEach((apartamento: any) => this.apartamentos.push(apartamento));
        },
        error => { }
      );
  }

  buscarApartamentoID(apartamento: ApartamentoModel): void {
    this.apartamentosService.buscarApartamentoID(apartamento.id!)
      .subscribe(
        response => {
          this.apartamento_id = response.id;
          this.apartamentoModalForm.patchValue({ identificador_modal: response.identificador, descricao_modal: response.descricao });
        },
        error => { }
      );
  }

  deletarApartamento(apartamento: ApartamentoModel): void {
    if (confirm(`Deseja deletar o apartamento ${apartamento.identificador}?`)) {
      this.apartamentosService.deletarApartamento(apartamento.id!)
        .subscribe(
          response => {
            this.buscarTodosApartamentos();
          },
          error => { }
        );
    }
  }

  atualizarApartamento(): void {
    this.apt_selec = { identificador: this.identificador_modal.value, descricao: this.descricao_modal.value };

    this.apartamentosService.atualizarApartamento(this.apartamento_id, this.apt_selec)
      .subscribe(
        response => {
          this.buscarTodosApartamentos();
        },
        error => { }
      );
  }

}
