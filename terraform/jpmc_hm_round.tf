variable "aws-reagin"{
  default = "us-east-1"
}

porvider "aws"{
  region = var.aws_region
}

resource "aws_security_group" "exmaple-security"{
  name = "some_securty"

  ingress {
    from_port = 80
    to_prot = 80
    protocol = "tcp"
    cidr_blocks = [0.0.0.0/0]
  }

  egress{
    from_port = 80
    to_prot = 80
    protocol = "tcp"
    cidr_blocks = [0.0.0.0/0]
  }
}

resouce "aws_instance" "example_ecinstance"{
  ami 
  name = 
  type = 

  vpc_security_groud_id = [aws_security_group.exmaple-security.id]

  tag = {

  }
}