variable "aws_region" {
  default = "us-west-2"
}

variable "instance_type" {
  default = "t2.micro"
}

variable "ami_id" {
  description = "AMI ID for the EC2 instance"
  default     = "ami-0c55b159cbfafe1f0"
}

variable "key_name" {
  description = "SSH key name"
  default     = "my-key-pair"
}

variable "allowed_ports" {
  type    = list(number)
  default = [22, 80, 443]
}

provider "aws" {
  region = var.aws_region
}

resource "aws_security_group" "example" {
  name        = "example-security-group"
  description = "Allow SSH, HTTP, and HTTPS traffic"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "example" {
  ami           = var.ami_id
  instance_type = var.instance_type
  key_name      = var.key_name

  vpc_security_group_ids = [aws_security_group.example.id]

  tags = {
    Name = "ExampleInstance"
  }
}

output "instance_id" {
  description = "The ID of the EC2 instance"
  value       = aws_instance.example.id
}

output "instance_public_ip" {
  description = "The public IP address of the EC2 instance"
  value       = aws_instance.example.public_ip
}

output "security_group_id" {
  description = "The ID of the security group"
  value       = aws_security_group.example.id
}
