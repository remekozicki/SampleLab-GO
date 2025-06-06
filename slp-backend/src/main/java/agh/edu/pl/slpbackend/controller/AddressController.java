package agh.edu.pl.slpbackend.controller;

import agh.edu.pl.slpbackend.controller.iface.AbstractController;
import agh.edu.pl.slpbackend.dto.AddressDto;
import agh.edu.pl.slpbackend.service.AddressService;
import lombok.AllArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@AllArgsConstructor
@RequestMapping("/address")
@CrossOrigin(origins = "http://localhost:3000")
public class AddressController extends AbstractController {

    private final AddressService addressService;

    @GetMapping("/list")
    public ResponseEntity<List<AddressDto>> list() {
        return ResponseEntity.ok(addressService.selectAll());
    }
}
