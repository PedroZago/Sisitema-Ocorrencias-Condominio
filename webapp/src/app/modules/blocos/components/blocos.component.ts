import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';

import { BlocosService } from '../services';
import { BlocoModel } from '../models';

@Component({
  selector: 'app-blocos',
  templateUrl: './blocos.component.html',
  styleUrls: ['./blocos.component.css']
})
export class BlocosComponent implements OnInit {

  bloco!: BlocoModel;
  blocos: any = [];
  blc_selec: BlocoModel = {};
  bloco_id!: number;

  constructor(
    private blocosService: BlocosService
  ) { }

  ngOnInit(): void {
    this.buscarTodosBlocos();
  }

  blocoForm = new FormGroup({
    identificador: new FormControl(null, [
      Validators.required
    ]),
    descricao: new FormControl(null, [
      Validators.required
    ])
  });

  blocoModalForm = new FormGroup({
    identificador_modal: new FormControl(null, [
      Validators.required
    ]),
    descricao_modal: new FormControl(null, [
      Validators.required
    ])
  });

  get identificador(): any {
    return this.blocoForm.get('identificador');
  }

  get descricao(): any {
    return this.blocoForm.get('descricao');
  }

  get identificador_modal(): any {
    return this.blocoModalForm.get('identificador_modal');
  }

  get descricao_modal(): any {
    return this.blocoModalForm.get('descricao_modal');
  }

  cadastrarBloco(): void {
    this.bloco = { descricao: this.descricao.value, identificador: this.identificador.value };

    this.blocosService.cadastrarBloco(this.bloco)
      .subscribe(
        response => {
          this.buscarTodosBlocos();
          this.blocoForm.reset();
        },
        error => { }
      );
  }

  buscarTodosBlocos(): void {
    this.blocosService.buscarTodosBlocos()
      .subscribe(
        response => {
          this.blocos = [];
          response.forEach((bloco: any) => this.blocos.push(bloco));
        },
        error => { }
      );
  }

  buscarBlocoID(apartamento: BlocoModel): void {
    this.blocosService.buscarBlocoID(apartamento.id!)
      .subscribe(
        response => {
          this.bloco_id = response.id;
          this.blocoModalForm.patchValue({ identificador_modal: response.identificador, descricao_modal: response.descricao });
        },
        error => { }
      );
  }

  deletarBloco(bloco: BlocoModel): void {
    if (confirm(`Deseja deletar o bloco ${bloco.identificador}?`)) {
      this.blocosService.deletarBloco(bloco.id!)
        .subscribe(
          response => {
            this.buscarTodosBlocos();
          },
          error => { }
        );
    }
  }

  atualizarBloco(): void {
    this.blc_selec = { identificador: this.identificador_modal.value, descricao: this.descricao_modal.value };

    this.blocosService.atualizarBloco(this.bloco_id, this.blc_selec)
      .subscribe(
        response => {
          this.buscarTodosBlocos();
        },
        error => { }
      );
  }

}
